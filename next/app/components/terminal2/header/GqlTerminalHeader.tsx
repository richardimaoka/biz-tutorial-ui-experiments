import { TerminalHeader } from "./TerminalHeader";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlTerminalHeader on Terminal2 {
    currentDirectory
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlTerminalHeader(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const currentDirectory = fragment.currentDirectory
    ? fragment.currentDirectory
    : "";
  return <TerminalHeader currentDirectory={currentDirectory} selectTab="" />;
}
