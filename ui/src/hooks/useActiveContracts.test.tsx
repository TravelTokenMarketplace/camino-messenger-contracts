import { describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import { useActiveContracts } from "./useActiveContracts";

vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));

describe("useActiveContracts", () => {
  it("reports supported=true for a chain with contracts", () => {
    const { result } = renderHook(() => useActiveContracts());
    expect(result.current.chainId).toBe(84532);
    expect(result.current.supported).toBe(true);
    expect(result.current.manager).toBeTruthy();
  });
});
