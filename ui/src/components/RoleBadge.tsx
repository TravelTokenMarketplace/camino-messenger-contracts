export function RoleBadge({ role }: { role: string }) {
  return (
    <span className="inline-block shrink-0 whitespace-nowrap rounded-[3px] border border-tarmac-200 bg-tarmac-50 px-1.5 py-0.5 font-mono text-[0.625rem] uppercase tracking-[0.08em] text-tarmac-600 dark:border-tarmac-700 dark:bg-tarmac-800 dark:text-tarmac-300">
      {role}
    </span>
  );
}
