// Global Declarations
export type Record<K extends string | number | symbol, T> = { [P in K]: T };

export type Nullable<T> = T | null;

//
// namespace tstest
//

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

  export type Season = typeof EnumSeason[keyof typeof EnumSeason];

  export type Test = typeof EnumTest[keyof typeof EnumTest];

  export type MyType = number;

  export type TestTypeStruct = TestStruct1;

  export type TestTypeTime = Date;

  export type TestType = number[];

  export type TestTypeMap = Record<string, Record<number, string>>;

  export type Direction = typeof EnumDirection[keyof typeof EnumDirection];

  export const EnumDirection = {
    North: 0,
    East: 1,
    South: 2,
    West: 3,
  } as const;

  export const EnumSeason = {
    Summer: "summer",
    Autumn: "autumn",
    Winter: "winter",
    Spring: "spring",
  } as const;

  export const EnumTest = {
    A: 0,
    B: 1,
    C: 2,
    D: 3,
  } as const;

  export const Timeout = 1000;

  export const Uno = "uno";

  export const Cento = 100;
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
