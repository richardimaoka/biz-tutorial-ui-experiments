import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { TerminalComponent } from "../terminal/TerminalComponent";
import { ColumnContentsPosition } from "./ColumnContentsPosition";

const fragmentDefinition = graphql(`
  fragment TerminalColumnFragment on TerminalColumn {
    terminal {
      ...TerminalComponent_Fragment
    }
  }
`);

export interface TerminalColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const TerminalColumn = (props: TerminalColumnProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.terminal) {
    return <></>;
  }

  return <TerminalComponent fragment={fragment.terminal} />;
};
