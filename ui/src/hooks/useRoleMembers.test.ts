import { describe, expect, it } from "vitest";
import { toMemberList } from "./useRoleMembers";

describe("toMemberList", () => {
  it("returns [] for undefined", () => {
    expect(toMemberList(undefined)).toEqual([]);
  });

  it("stringifies address array entries", () => {
    expect(toMemberList(["0xAAA", "0xBBB"])).toEqual(["0xAAA", "0xBBB"]);
  });
});
