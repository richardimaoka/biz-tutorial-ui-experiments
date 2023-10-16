import { CommandAnimation } from "@/app/components/terminal2/CommandAnimation";

export default async function Page() {
  return (
    <div>
      <CommandAnimation command="npx live-server" />
    </div>
  );
}
