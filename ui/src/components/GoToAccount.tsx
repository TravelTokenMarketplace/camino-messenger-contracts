import { useState } from "react";
import { ArrowRight } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { isAddress } from "viem";
import { Input } from "./Input";

export function GoToAccount() {
  const [value, setValue] = useState("");
  const navigate = useNavigate();
  const trimmed = value.trim();
  const valid = isAddress(trimmed);

  function go(e: React.FormEvent) {
    e.preventDefault();
    if (valid) navigate(`/account/${trimmed}`);
  }

  return (
    <form onSubmit={go} className="flex flex-col gap-1">
      <div className="flex items-center gap-2">
        <Input
          className="flex-1 font-mono"
          placeholder="0x… CM Account address"
          value={value}
          onChange={(e) => setValue(e.target.value)}
          aria-label="CM Account address"
        />
        <button
          type="submit"
          disabled={!valid}
          className="inline-flex items-center gap-1.5 rounded bg-indigo-600 px-3 py-1.5 text-white transition-colors hover:bg-indigo-700 disabled:opacity-50"
        >
          Go <ArrowRight className="h-4 w-4" />
        </button>
      </div>
      {trimmed.length > 0 && !valid && (
        <span className="text-xs text-red-600">Enter a valid address.</span>
      )}
    </form>
  );
}
