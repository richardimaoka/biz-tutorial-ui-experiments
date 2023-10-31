import { TerminalHeaderGql } from "@/app/components/terminal2/header/TerminalHeaderGql";
import { TerminalHeaderGqlFragment } from "@/libs/gql/graphql";

export default async function Page() {
  const fragment: TerminalHeaderGqlFragment = {
    __typename: "Terminal2",
    currentDirectory: "dirdir",
  };

  return <div>{<TerminalHeaderGql fragment={fragment as any} />}</div>;
}
