import { describe, expect, it } from "vitest";
import { encodeEventTopics, parseAbi } from "viem";
import { findCreatedAccount } from "./receipt";

const abi = parseAbi(["event CMAccountCreated(address indexed account)"]);

describe("findCreatedAccount", () => {
  it("extracts the account address from logs", () => {
    const addr = "0x1111111111111111111111111111111111111111";
    // CMAccountCreated has only an indexed arg, so the value lives in topics; data is empty.
    const topics = encodeEventTopics({ abi, eventName: "CMAccountCreated", args: { account: addr } });
    const found = findCreatedAccount([{ data: "0x", topics }] as never, abi as never);
    expect(found?.toLowerCase()).toBe(addr);
  });

  it("returns undefined when no matching event", () => {
    expect(findCreatedAccount([{ data: "0x", topics: ["0xdead"] }] as never, abi as never)).toBeUndefined();
  });
});
