import { describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import { useActiveContracts } from "./useActiveContracts";

vi.mock("wagmi", () => ({ useChainId: () => 84532 }));

describe("useActiveContracts", () => {
  it("reports supported=true for a chain with contracts", () => {
    const { result } = renderHook(() => useActiveContracts());
    expect(result.current.chainId).toBe(84532);
    expect(result.current.supported).toBe(true);
    expect(result.current.manager).toBeTruthy();
  });
});
