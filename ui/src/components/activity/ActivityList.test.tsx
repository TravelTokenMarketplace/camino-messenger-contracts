import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { ActivityList } from "./ActivityList";

describe("ActivityList", () => {
  it("shows the Load older button even when the window is empty", () => {
    const onLoadOlder = vi.fn();
    render(
      <ActivityList events={[]} isLoading={false} hasNextPage onLoadOlder={onLoadOlder} oldestBlockLoaded={50_000n} />,
    );

    expect(screen.getByText("No activity in the last 10,000 blocks.")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Load older" })).toBeEnabled();
    expect(screen.getByText("Scanned back to block 50000")).toBeInTheDocument();
  });

  it("omits the Load older button entirely on the Dashboard (no onLoadOlder)", () => {
    render(<ActivityList events={[]} isLoading={false} />);
    expect(screen.getByText("No activity in the last 10,000 blocks.")).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Load older" })).not.toBeInTheDocument();
  });

  it("disables the button when there is no more history", () => {
    render(<ActivityList events={[]} isLoading={false} hasNextPage={false} onLoadOlder={vi.fn()} />);
    expect(screen.getByRole("button", { name: "No more history" })).toBeDisabled();
  });
});
