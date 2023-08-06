import { TerminalCommandComponent } from "./TerminalCommandComponent";
import { TerminalOutputComponent } from "./TerminalOutputComponent";
import styles from "./terminal.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalNodeComponent_Fragment on TerminalNode {
    content {
      __typename
      ... on TerminalCommand {
        ...TerminalCommand_Fragment
      }
      ... on TerminalOutput {
        ...TerminalOutput_Fragment
      }
    }
  }
`);

export interface TerminalNodeComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const TerminalNodeComponent = (
  props: TerminalNodeComponentProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment?.content) {
    return <></>;
  }

  switch (fragment.content.__typename) {
    case "TerminalCommand":
      return <TerminalCommandComponent fragment={fragment.content} />;
    case "TerminalOutput":
      return <TerminalOutputComponent fragment={fragment.content} />;
  }
};
