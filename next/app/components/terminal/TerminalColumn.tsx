import { TerminalComponent } from "./TerminalComponent";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalColumn_Fragment on TerminalColumn {
    terminal {
      ...TerminalComponent_Fragment
    }
  }
`);

interface TerminalColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  skipAnimation?: boolean;
  isFocused: boolean;
}

export const TerminalColumn = (props: TerminalColumnProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.terminal) {
    return <></>;
  }

  return (
    <TerminalComponent
      fragment={fragment.terminal}
      skipAnimation={props.skipAnimation}
      isFocused={props.isFocused}
    />
  );
};
