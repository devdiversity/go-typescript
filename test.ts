type Nullable<T> = T | null;
type MyType = number

export namespace tsTestExternal {
  export interface UserRegisterResponse {
    token: string;
    user?: string;
  }

  export interface UserRegisterResponse2 {
    token: string;
    user?: string;
    testdep: tsTestExternal2.TestExternal2;
  }
}
// end namespace tsTestExternal

export namespace tsTestExternal2 {
  export interface TestExternal2 {
    token: string;
    user?: string;
  }
}
// end namespace tsTestExternal2

export namespace tstest {
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
    MapsNested: Record<string, Record<number, string>>;
    MapsNestedPtr: Record<string, Record<number, Nullable<string[]>>>;
    testTest: TestTest;
    testType: TestType;
    TestTypeMap: TestTypeMap;
    testdep: tsTestExternal.UserRegisterResponse2;
  }

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

  export type TestType = number[];

  export type TestTypeMap = Record<string, Record<number, string>>;
}
// end namespace tstest
