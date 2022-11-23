
// Api Class
export interface ApiRestResponse {
  data?: unknown;
  error: string | null;
}

class Api {
  apiUrl: string;
  localStorage: Storage | null;

  constructor(apiurl: string) {
    this.apiUrl = apiurl;
    this.localStorage = null;
  }

  request(method: string, url: string, data: unknown, timeout = 7000, upload = false) {
    return new Promise((resolve, reject) => {
      let auth: string;
      const headers: { [key: string]: string } = {
        "Content-Type": upload ? "multipart/form-data" : "application/json",
        "Cache-Control": "no-cache",
      };
      if (this.localStorage) {
        auth = localStorage.getItem("jwt-token") as string;
        headers["auth-token"] = auth;
      }
      const controller = new AbortController();
      const timeoutId = setTimeout(() => controller.abort(), timeout);

      let requestOptions: RequestInit = {
        method: method,
        body: data ? JSON.stringify(data) : undefined,
        cache: "no-store",
        mode: "cors",
        credentials: "include",
        headers: headers || {},
        signal: controller.signal,
      };

      if (upload) {
        requestOptions = {
          method: method,
          body: data as FormData,
          signal: controller.signal,
        };
      }

      fetch(url, requestOptions)
        .then((response) => {
          if (!response.ok) {
            const error = `api.error.${response.statusText}`;
            throw error;
          } else {
            if (this.localStorage) {
              const jwt = response.headers.get("auth-token");
              if (jwt) {
                this.localStorage.setItem("jwt-token", jwt);
              }
            }
            return response.json() as Promise<ApiRestResponse>;
          }
        })
        .then((data) => {
          clearTimeout(timeoutId);
          if (typeof data === "object" && data.hasOwnProperty("data") && data.hasOwnProperty("error")) {
            const d: ApiRestResponse = data;
            if (d.error) {
              throw d.error;
            } else {
              resolve(d);
            }
          } else {
            throw "api.error.wrongdatatype";
          }
        })
        .catch((error) => {
          clearTimeout(timeoutId);
          if ((error as Error).toString() === "DOMException: The user aborted a request.") {
            reject(new Error("api.error.timeouterror"));
            return;
          }
          if ((error as Error).toString() === "TypeError: Failed to fetch") {
            reject(new Error("api.error.connectionerror"));
            return;
          }
          reject(error);
        });
    });
  }

  processResult(url: string, result: ApiRestResponse): { data: unknown; error: string | null } {
    if (typeof result.data !== "object") {
      return { data: result.data, error: null };
    } else if (!result.data) {
      result.data = {};
    }
    return { data: result.data, error: null };
  }

  processError(
    error: Error,
    url: string
  ): {
    data: unknown;
    error: string | null;
  } {
    if (error.message === "api.error.timeouterror") {
      Object.defineProperty(error, "__api_error__", {
        value: error.message,
        writable: false,
      });

      return { data: null, error: error.message };
    }
    if (error.message === "api.error.connectionerror") {
      Object.defineProperty(error, "__api_error__", {
        value: error.message,
        writable: false,
      });

      return { data: null, error: error.message };
    }

    return {
      data: null,
      error: error.message,
    };
  }

  async POST(
    url: string,

    data: unknown,
    timeout?: number
  ): Promise<{
    data: unknown;
    error: string | null;
  }> {
    try {
      let upload = false;
      if (url.includes("/upload/")) {
        upload = true;
      }
      const result = (await this.request("POST", `${this.apiUrl}${url}`, data, timeout, upload)) as ApiRestResponse;
      return this.processResult(url, result);
    } catch (error) {
      return new Promise<{
        data: unknown;
        error: string | null;
      }>(async (resolve) => {
        let result = this.processError(error, `POST => ${this.apiUrl}${url}`);
        resolve(result);
      });
    }
  }

  async GET(
    url: string,
    timeout?: number
  ): Promise<{
    data: unknown;
    error: string | null;
  }> {
    try {
      const result = (await this.request("GET", `${this.apiUrl}${url}`, null, timeout)) as ApiRestResponse;
      return this.processResult(url, result);
    } catch (error) {
      return new Promise<{
        data: unknown;
        error: string | null;
      }>(async (resolve) => {
        let result = this.processError(error, `GET => ${this.apiUrl}${url}`);
        resolve(result);
      });
    }
  }

  async UPLOAD(
    url: string,
    data: unknown,
    timeout?: number
  ): Promise<{
    data: unknown;
    error: string | null;
  }> {
    try {
      const result = (await this.request("POST", `${this.apiUrl}${url}`, data, timeout, true)) as ApiRestResponse;
      return this.processResult(url, result);
    } catch (error) {
      return new Promise<{
        data: unknown;
        error: string | null;
      }>((resolve) => {
        resolve(this.processError(error, `POST => ${this.apiUrl}${url}`));
      });
    }
  }
}

const api = new Api("http://localhost:8080");

// Global Declarations 
export type Nullable<T> = T | null;

export type Record<K extends string | number | symbol, T> = { [P in K]: T; }


//
// namespace server
//

export namespace server {

export interface FormRequest {
	req: string;
	count: number;
}


export interface FormResponse {
	test: string;
}


export interface HTTPResponse {
	data: unknown;
	error: unknown;
}


// Typescript: TSEndpoint= path=/postTest;  name=postTest; method=POST; request=FormRequest; response=FormResponse
// server/server.go Line: 70
export const postTest = async (data: FormRequest):Promise<{ data:FormResponse; error: Nullable<string> }> => {
	return await api.POST("/postTest", data) as { data: FormResponse; error: Nullable<string> };
}

// Typescript: TSEndpoint= path=/ping; name=ping; method=GET; response=string
// server/server.go Line: 50
export const ping = async ():Promise<{ data:string; error: Nullable<string> }> => {
	return await api.GET("/ping") as { data: string; error: Nullable<string> };
}

// Typescript: TSEndpoint= path=/test/:len/*lat;  name=test; method=GET; response=string
// server/server.go Line: 56
export const test = async (len: string, lat: Nullable<string>):Promise<{ data:string; error: Nullable<string> }> => {
	return await api.GET(`/test/${len}/${lat}`) as { data: string; error: Nullable<string> };
}

// Typescript: TSEndpoint= path=/testpost;  name=testPost; method=POST; request=FormRequest; response=string
// server/server.go Line: 63
export const testPost = async (data: FormRequest):Promise<{ data:string; error: Nullable<string> }> => {
	return await api.POST("/testpost", data) as { data: string; error: Nullable<string> };
}
}


//
// namespace tstest
//

export namespace tstest {

export interface HttpSessions {
	id: number;
	key: string;
	data: string;
	created: Date;
	modified: Date;
	expire: Date;
}


export interface TestTest {
	info: string;
	typename: MyType;
	created: Date;
}


export interface TestStruct1 {
	created: Date;
	TestT: TestType;
	Session: HttpSessions;
	id: number;
	key: string[];
	data: Nullable<string>;
	newpassword: tsTestExternal.UserRegisterResponse;
	modified: Date;
	maps: Date;
	MapsNested: Record<string , Record<number , string>>;
	MapsNestedPtr: Record<string , Record<number , Nullable<string[]>>>;
	testTest: TestTest;
	testType: TestType;
	TestTypeMap: TestTypeMap;
	testdep: tsTestExternal.UserRegisterResponse2;
	direction: Direction;
	season: Season;
}

export type Direction = typeof EnumDirection[keyof typeof EnumDirection] 

export type Season = typeof EnumSeason[keyof typeof EnumSeason] 

export type TestTypeStruct = TestStruct1

export type TestTypeTime = Date

export type Test = typeof EnumTest[keyof typeof EnumTest] 

export type MyType = number

export type TestType = number[]

export type TestTypeMap = Record<string , Record<number , string>>

export const EnumSeason = {
Summer: "summer",
Autumn: "autumn",
Winter: "winter",
Spring: "spring",
} as const

export const EnumTest = {
A: 0,
B: 1,
C: 2,
D: 3,
} as const

export const EnumDirection = {
North: 0,
East: 1,
South: 2,
West: 3,
} as const

export const Timeout = 1000

export const Uno = "uno"

export const Cento = 100

}


//
// namespace tsTestExternal
//

export namespace tsTestExternal {

export interface UserRegisterResponse {
	token: string;
	user?: string;
}


export interface UserRegisterResponse2 {
	token: string;
	user?: string;
	testdep: moduleExt.ModuleExtTest;
}

}


//
// namespace moduleExt
//

export namespace moduleExt {

export interface ModuleExtTest {
	token: string;
	user?: Nullable<string>;
}

}

