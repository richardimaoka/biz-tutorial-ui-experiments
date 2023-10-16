import { CommandTypeIn } from "@/app/components/terminal2/CommandTypeIn";

export default async function Page() {
  return (
    <div>
      <CommandTypeIn command="npx live-server" />
    </div>
  );
}
