import { TextDecoder, TextEncoder } from "util";
import {
  Server,
  credentials,
  ServerUnaryCall,
  sendUnaryData,
} from "@grpc/grpc-js";
import { v4 as uuidv4 } from "uuid";

import {
  HttpMethod as ProtoHttpMethod,
  HttpMethodMap as ProtoHttpMethodMap,
  Manifest,
  Route,
  Version as ProtoVersion,
} from "../dist/proto/manifest_pb";
import {
  FoldIngressService,
  IFoldIngressServer,
} from "../dist/proto/ingress_grpc_pb";
import {
  ManifestReq,
  Request as ProtoRequest,
  Response as ProtoResponse,
  StringArray,
} from "../dist/proto/ingress_pb";
import { Map as ProtoMap } from "google-protobuf";

import { getLogger, Logger } from "./logging";

export enum HttpMethod {
  GET = "GET",
  HEAD = "HEAD",
  POST = "POST",
  PUT = "PUT",
  DELETE = "DELETE",
  CONNECT = "CONNECT",
  OPTIONS = "OPTIONS",
  TRACE = "TRACE",
  PATCH = "PATCH",
}

export class Request {
  public httpMethod!: HttpMethod;
  public path!: string;
  public handler!: string;
  public body!: { [key: string]: any };
  public headers!: { [key: string]: string[] };
  public pathParams!: { [key: string]: string };
  public queryParams!: { [key: string]: string[] };
}

export class Response {
  public statusCode!: number;
  public body!: { [key: string]: any };
  public headers!: { [key: string]: string[] };
}

export type Handler = (req: Request, res: Response) => void;

export interface Version {
  major: number;
  minor: number;
  patch: number;
}

export class Service {
  private name: string;
  private grpcBackend: GrpcBackend;
  private _logger: Logger;

  constructor() {
    this.name = process.env.FOLD_SERVICE_NAME!;
    this._logger = getLogger(this.name);
    this.grpcBackend = new GrpcBackend(this.name, this._logger);
  }

  public set version(version: Version) {
    this.grpcBackend.version = version;
  }

  public get logger(): Logger {
    return this._logger;
  }

  public set logger(logger: Logger) {
    this._logger = logger;
    this.grpcBackend.logger = logger;
  }

  public start(): void {
    this.grpcBackend.serve();
  }

  public get(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.GET, path, handler);
  }

  public head(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.HEAD, path, handler);
  }

  public post(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.POST, path, handler);
  }

  public put(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.PUT, path, handler);
  }

  public delete(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.DELETE, path, handler);
  }

  public connect(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.CONNECT, path, handler);
  }

  public options(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.OPTIONS, path, handler);
  }

  public trace(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.TRACE, path, handler);
  }

  public patch(path: string, handler: Handler): void {
    this.grpcBackend.registerHandler(HttpMethod.PATCH, path, handler);
  }
}

class GrpcBackend {
  private handlers: { [key: string]: Handler };
  private _manifest: Manifest;
  private server!: Server;
  private _logger: Logger;

  constructor(name: string, logger: Logger) {
    this.handlers = {};
    this._manifest = new Manifest();
    this._manifest.setName(name);
    this._logger = logger;
  }

  get manifest(): Manifest {
    return this._manifest;
  }

  get logger(): Logger {
    return this._logger;
  }

  set logger(logger: Logger) {
    this._logger = logger;
  }

  set version(version: Version) {
    let v: ProtoVersion = new ProtoVersion();
    v.setMajor(version.major);
    v.setMinor(version.minor);
    v.setPatch(version.patch);
    this.manifest.setVersion(v);
  }

  registerHandler(method: HttpMethod, path: string, handler: Handler): void {
    this.logger.debug(`registering handler ${method} ${path}`);
    let handlerId: string = uuidv4();
    let route = new Route();
    route.setHttpMethod(ProtoHttpMethod[HttpMethod[method]]);
    route.setHandler(handlerId);
    route.setPathSpec(path);
    this.manifest.addRoutes(route);
    this.handlers[handlerId] = handler;
  }

  handleRequest(request: Request): Response {
    let response: Response = new Response();
    const handler: Handler = this.handlers[request.handler];
    handler(request, response);
    return response;
  }

  serve(): void {
    const socketAddr = `unix://${process.env.FOLD_SOCK_ADDR!}`;
    this.logger.debug(`starting server on socket ${socketAddr}`);
    this.server = new Server();
    this.server.addService(FoldIngressService, newFoldIngressServer(this));
    this.server.bindAsync(
      socketAddr,
      credentials.createInsecure() as any,
      (err: Error | null, bindPort: number) => {
        if (err) {
          throw err;
        }
        this.logger.debug(
          `binding server to socket ${socketAddr} with port ${bindPort}`
        );
        this.server.start();
      }
    );
  }
}

function newFoldIngressServer(backend: GrpcBackend): IFoldIngressServer {
  return {
    getManifest(
      _: ServerUnaryCall<ManifestReq, Manifest>,
      callback: sendUnaryData<Manifest>
    ): void {
      backend.logger.debug(`retrieving service manifest`);
      callback(null, backend.manifest);
    },

    doRequest(
      call: ServerUnaryCall<ProtoRequest, ProtoResponse>,
      callback: sendUnaryData<ProtoResponse>
    ): void {
      let request = decodeProtoRequest(call.request);
      backend.logger.debug(
        `handling request for ${request.httpMethod} ${request.path}`
      );
      const response = backend.handleRequest(request);
      callback(null, encodeProtoResponse(response));
    },
  };
}

function decodeProtoHttpMethod(
  n: ProtoHttpMethodMap[keyof ProtoHttpMethodMap]
): HttpMethod {
  switch (n) {
    case ProtoHttpMethod.GET:
      return HttpMethod.GET;

    case ProtoHttpMethod.HEAD:
      return HttpMethod.HEAD;

    case ProtoHttpMethod.POST:
      return HttpMethod.POST;

    case ProtoHttpMethod.PUT:
      return HttpMethod.PUT;

    case ProtoHttpMethod.DELETE:
      return HttpMethod.DELETE;

    case ProtoHttpMethod.CONNECT:
      return HttpMethod.CONNECT;

    case ProtoHttpMethod.OPTIONS:
      return HttpMethod.OPTIONS;

    case ProtoHttpMethod.TRACE:
      return HttpMethod.TRACE;

    case ProtoHttpMethod.PATCH:
      return HttpMethod.PATCH;

    default:
      // This can't happen actually happen in practice as it has all been
      // validated by the gRPC server.
      throw new Error(`Invalid HTTP method ${n}`);
  }
}

function decodeProtoRequest(req: ProtoRequest): Request {
  let request = new Request();
  request.httpMethod = decodeProtoHttpMethod(req.getHttpMethod());
  request.path = req.getPath();
  request.handler = req.getHandler();
  if (
    request.httpMethod == HttpMethod.GET ||
    request.httpMethod == HttpMethod.DELETE
  ) {
    request.body = {};
  } else {
    request.body = JSON.parse(
      new TextDecoder("utf-8").decode(req.getBody_asU8())
    );
  }
  request.headers = decodeMapStringArray(req.getHeadersMap());
  request.pathParams = decodeMapString(req.getPathParamsMap());
  request.queryParams = decodeMapStringArray(req.getQueryParamsMap());
  return request;
}

function encodeProtoResponse(res: Response): ProtoResponse {
  let response = new ProtoResponse();
  response.setStatus(res.statusCode);
  response.setBody(new TextEncoder().encode(JSON.stringify(res.body)));
  return response;
}

function decodeMapStringArray(
  map: ProtoMap<string, StringArray>
): { [key: string]: string[] } {
  let results: { [key: string]: string[] } = {};
  map.forEach((value: StringArray, key: string) => {
    results[key] = value.getValuesList();
  });
  return results;
}

function decodeMapString(
  map: ProtoMap<string, string>
): { [key: string]: string } {
  let results: { [key: string]: string } = {};
  map.forEach((value: string, key: string) => {
    results[key] = value;
  });
  return results;
}
