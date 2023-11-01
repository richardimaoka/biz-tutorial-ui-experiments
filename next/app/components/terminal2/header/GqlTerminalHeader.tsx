import { TerminalHeader } from "./TerminalHeader";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlTerminalHeader on TerminalColumn2 {
    terminals {
      name
      currentDirectory
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectIndex: number;
}

export function GqlTerminalHeader(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const selectTerminal = fragment.terminals[props.selectIndex];

  const currentDirectory = selectTerminal.currentDirectory
    ? selectTerminal.currentDirectory
    : "";

  const selectTab = selectTerminal.name ? selectTerminal.name : "";

  return (
    <TerminalHeader currentDirectory={currentDirectory} selectTab={selectTab} />
  );
}
