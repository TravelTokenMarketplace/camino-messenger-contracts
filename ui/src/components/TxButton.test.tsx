import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { TxButton } from "./TxButton";

describe("TxButton", () => {
  it("calls write and shows confirmed on success", async () => {
    const write = vi.fn().mockResolvedValue("0xhash");
    const onConfirmed = vi.fn();
    render(<TxButton label="Do it" write={write} onConfirmed={onConfirmed} />);
    fireEvent.click(screen.getByRole("button", { name: /do it/i }));
    await waitFor(() => expect(write).toHaveBeenCalled());
    await waitFor(() => expect(onConfirmed).toHaveBeenCalled());
  });

  it("shows an error message on failure", async () => {
    const write = vi.fn().mockRejectedValue(new Error("user rejected"));
    render(<TxButton label="Do it" write={write} />);
    fireEvent.click(screen.getByRole("button", { name: /do it/i }));
    await waitFor(() =>
      expect(screen.getByText(/user rejected/i)).toBeInTheDocument(),
    );
  });
});
