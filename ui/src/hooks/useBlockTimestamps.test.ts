import { describe, expect, it } from "vitest";
import { dedupeBlockNumbers } from "./useBlockTimestamps";

describe("dedupeBlockNumbers", () => {
  it("removes duplicates while preserving first-seen order", () => {
    expect(dedupeBlockNumbers([5n, 3n, 5n, 1n, 3n])).toEqual([5n, 3n, 1n]);
  });

  it("returns an empty array unchanged", () => {
    expect(dedupeBlockNumbers([])).toEqual([]);
  });
});
