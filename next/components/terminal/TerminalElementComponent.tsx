import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";

const TerminalElementComponent_Fragment = graphql(`
  fragment TerminalElementComponent_Fragment on TerminalElement {
    __typename
    ... on TerminalCommand {
      ...TerminalCommand_Fragment
    }
    ... on TerminalOutput {
      ...TerminalOutput_Fragment
    }
  }
`);

export interface TerminalElementComponentProps {
  fragment: FragmentType<typeof TerminalElementComponent_Fragment>;
}

export const TerminalElementComponent = (
  props: TerminalElementComponentProps
): JSX.Element => {
  const fragment = useFragment(
    TerminalElementComponent_Fragment,
    props.fragment
  );

  switch (fragment.__typename) {
    case "TerminalCommand":
      return <div>{fragment.__typename}</div>;
    case "TerminalOutput":
      return <div>{fragment.__typename}</div>;
    case "TerminalCommandSet":
      return <></>;
  }
};
