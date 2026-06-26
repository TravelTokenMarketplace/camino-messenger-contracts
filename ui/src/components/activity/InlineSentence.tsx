import { Fragment, type ReactNode } from "react";
import { shortAddress } from "../../lib/format";
import { CopyButton } from "../CopyButton";

/** Matches the `0xXXXX…XXXX` short form produced by shortAddress(). */
const SHORT_TOKEN = /0x[0-9a-fA-F]{4}…[0-9a-fA-F]{4}/g;
/** A full 0x value worth resolving (address = 40 hex, bytes32 = 64 hex). */
const FULL_HEX = /^0x[0-9a-fA-F]{8,}$/;

/**
 * Renders an activity sentence, upgrading any shortened `0x…` token back into an
 * interactive chip: hovering shows the full value and a copy button sits beside
 * it. The full values are recovered from the event's decoded `args`, so the
 * sentence string itself stays plain (it's still what tests and filters see).
 */
export function InlineSentence({ sentence, args }: { sentence: string; args: Record<string, unknown> }) {
  const fullByShort = new Map<string, string>();
  for (const v of Object.values(args)) {
    if (typeof v === "string" && FULL_HEX.test(v)) fullByShort.set(shortAddress(v), v);
  }

  const parts: ReactNode[] = [];
  let last = 0;
  let key = 0;
  SHORT_TOKEN.lastIndex = 0;
  for (let m = SHORT_TOKEN.exec(sentence); m !== null; m = SHORT_TOKEN.exec(sentence)) {
    if (m.index > last) parts.push(<Fragment key={key++}>{sentence.slice(last, m.index)}</Fragment>);
    const full = fullByShort.get(m[0]);
    if (full) {
      parts.push(
        <span key={key++} className="inline-flex items-center gap-1 whitespace-nowrap font-mono" title={full}>
          {m[0]}
          <CopyButton value={full} label="Copy address" />
        </span>,
      );
    } else {
      parts.push(<Fragment key={key++}>{m[0]}</Fragment>);
    }
    last = m.index + m[0].length;
  }
  if (last < sentence.length) parts.push(<Fragment key={key++}>{sentence.slice(last)}</Fragment>);

  return <>{parts}</>;
}
