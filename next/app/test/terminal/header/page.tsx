import { GqlTerminalHeader } from "@/app/components/terminal2/header/GqlTerminalHeader";
import { GqlTerminalHeaderFragment } from "@/libs/gql/graphql";

export default async function Page() {
  const fragment: GqlTerminalHeaderFragment = {
    __typename: "Terminal2",
    currentDirectory: "dirdir",
  };

  return <div>{<GqlTerminalHeader fragment={fragment as any} />}</div>;
}
