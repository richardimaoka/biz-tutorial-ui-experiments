import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { TerminalCommandComponent } from "./TerminalCommandComponent";
import { TerminalOutputComponent } from "./TerminalOutputComponent";

const TerminalNodeComponent_Fragment = graphql(`
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
  fragment: FragmentType<typeof TerminalNodeComponent_Fragment>;
  isLastElement: boolean;
  scrollIntoView: () => void;
}

export const TerminalNodeComponent = (
  props: TerminalNodeComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalNodeComponent_Fragment, props.fragment);
  const scrollIntoView = props.scrollIntoView;

  if (fragment.content) {
    switch (fragment.content.__typename) {
      case "TerminalCommand":
        return (
          <TerminalCommandComponent
            fragment={fragment.content}
            scrollIntoView={scrollIntoView}
          />
        );
      case "TerminalOutput":
        return <TerminalOutputComponent fragment={fragment.content} />;
    }
  } else {
    return <></>;
  }
};
