import { TerminalHeaderGqlFragment } from "@/libs/gql/graphql";
import { TerminalHeader } from "./TerminalHeader";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalHeaderGql on Terminal2 {
    currentDirectory
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function TerminalHeaderGql(props: Props) {
  const fragment: TerminalHeaderGqlFragment = useFragment(
    fragmentDefinition,
    props.fragment
  );
  const currentDirectory = fragment.currentDirectory
    ? fragment.currentDirectory
    : "";
  return <TerminalHeader currentDirectory={currentDirectory} selectTab="" />;
}
