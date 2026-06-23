import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import App from "./App";
import { ThemeProvider } from "./theme/theme";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined, isConnected: false, chainId: undefined }),
  useChainId: () => 84532,
  useConnect: () => ({ connect: vi.fn(), connectors: [] }),
  useDisconnect: () => ({ disconnect: vi.fn() }),
  useSwitchChain: () => ({ switchChain: vi.fn() }),
  useReadContract: () => ({ data: undefined, isLoading: false }),
  useReadContracts: () => ({ data: undefined, isLoading: false }),
  usePublicClient: () => undefined,
}));

vi.mock("./wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
  ActiveChainProvider: ({ children }: { children: React.ReactNode }) => children,
}));

describe("App", () => {
  it("renders the header title and connect button", () => {
    render(
      <ThemeProvider>
        <MemoryRouter initialEntries={["/"]}>
          <App />
        </MemoryRouter>
      </ThemeProvider>,
    );
    expect(screen.getByText(/camino messenger/i)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /connect/i })).toBeInTheDocument();
  });
});
