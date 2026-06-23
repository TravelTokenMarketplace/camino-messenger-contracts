import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { AccountWorkspace } from "./AccountWorkspace";

vi.mock("wagmi", () => ({
  useChainId: () => 84532,
  useBalance: () => ({ data: { formatted: "1.0", symbol: "ETH" } }),
  useAccount: () => ({ address: undefined }),
  useReadContract: () => ({ data: undefined, isLoading: false }),
  usePublicClient: () => undefined,
}));

describe("AccountWorkspace", () => {
  it("renders the tab bar with the account address", () => {
    const addr = "0x1111111111111111111111111111111111111111";
    render(
      <MemoryRouter initialEntries={[`/account/${addr}`]}>
        <Routes><Route path="account/:address" element={<AccountWorkspace />} /></Routes>
      </MemoryRouter>,
    );
    expect(screen.getByRole("link", { name: /bots/i })).toBeInTheDocument();
    expect(screen.getByRole("heading", { level: 1 })).toHaveTextContent(addr.slice(0, 6));
  });
});
