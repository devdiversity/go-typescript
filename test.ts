export namespace tstest {
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
    MapsNested: Record<string, Record<number, string>>;
    MapsNestedPtr: Record<string, Record<number, Nullable<string[]>>>;
    testTest: TestTest;
    testType: TestType;
    TestTypeMap: TestTypeMap;
    testdep: tsTestExternal.UserRegisterResponse2;
    direction: Direction;
    season: Season;
  }

  export interface HttpSessions {
    id: number;
    key: string;
    data: string;
    created: Date;
    modified: Date;
    expire: Date;
  }

  export type TestTypeMap = Record<string, Record<number, string>>;

  export type Direction = typeof EnumDirection[keyof typeof EnumDirection];

  export type Season = typeof EnumSeason[keyof typeof EnumSeason];

  export type MyType = number;

  export type Nullable<T> = T | null;

  export type TestTypeStruct = TestStruct1;

  export type Test = typeof EnumTest[keyof typeof EnumTest];

  export type TestType = number[];

  export const EnumDirection = {
    North: 0,
    East: 1,
    South: 2,
    West: 3,
  } as const;

  export const EnumSeason = {
    Ss: "summer",
    As: "autumn",
    Ws: "winter",
    S2: "spring",
  } as const;

  export const EnumTest = {
    A: 0,
    B: 1,
    C: 2,
    D: 3,
  } as const;
}
// end namespace tstest

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
// end namespace tsTestExternal

export namespace moduleExt {
  export interface ModuleExtTest {
    token: string;
    user?: string;
  }
}
// end namespace moduleExt
