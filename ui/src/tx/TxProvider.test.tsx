import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { TxProvider, useTx } from "./TxProvider";

// wagmi's useConfig only needs to return an object the mocked actions ignore.
vi.mock("wagmi", () => ({ useConfig: () => ({}) }));

const waitMock = vi.fn();
vi.mock("wagmi/actions", () => ({
  getAccount: () => ({ chainId: 84532 }),
  waitForTransactionReceipt: (...args: unknown[]) => waitMock(...args),
}));

function Harness({ onConfirmed }: { onConfirmed: () => void }) {
  const { txs, track } = useTx();
  return (
    <div>
      <button onClick={() => void track({ label: "Do thing", write: async () => "0xhash", onConfirmed })}>go</button>
      {txs.map((t) => (
        <span key={t.id} data-testid="tx">{t.label}:{t.state}</span>
      ))}
    </div>
  );
}

function renderHarness(onConfirmed: () => void) {
  return render(
    <QueryClientProvider client={new QueryClient()}>
      <TxProvider>
        <Harness onConfirmed={onConfirmed} />
      </TxProvider>
    </QueryClientProvider>,
  );
}

describe("TxProvider", () => {
  it("tracks a tx and only confirms after the receipt is mined", async () => {
    let resolveReceipt!: (r: { status: string }) => void;
    waitMock.mockReturnValue(new Promise((res) => { resolveReceipt = res; }));
    const onConfirmed = vi.fn();

    renderHarness(onConfirmed);

    fireEvent.click(screen.getByText("go"));

    // Submitted: panel shows a pending entry, but onConfirmed has NOT fired yet.
    await waitFor(() => expect(screen.getByTestId("tx")).toHaveTextContent("Do thing:pending"));
    expect(onConfirmed).not.toHaveBeenCalled();

    // Mining completes successfully.
    resolveReceipt({ status: "success" });
    await waitFor(() => expect(screen.getByTestId("tx")).toHaveTextContent("Do thing:confirmed"));
    expect(onConfirmed).toHaveBeenCalledTimes(1);
  });

  it("stays confirmed even if onConfirmed throws", async () => {
    waitMock.mockResolvedValue({ status: "success" });
    const onConfirmed = vi.fn(() => {
      throw new Error("navigation blew up");
    });

    renderHarness(onConfirmed);
    fireEvent.click(screen.getByText("go"));

    // A throwing side effect must not flip a mined tx back to "failed".
    await waitFor(() => expect(screen.getByTestId("tx")).toHaveTextContent("Do thing:confirmed"));
    expect(onConfirmed).toHaveBeenCalledTimes(1);
  });
});
