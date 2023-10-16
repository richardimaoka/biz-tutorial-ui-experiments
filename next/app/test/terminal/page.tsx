import { CommandComponent } from "@/app/components/terminal2/CommandComponent";

export default async function Page() {
  return (
    <div>
      <CommandComponent command="npx live-server" />
    </div>
  );
}
