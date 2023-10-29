import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { CommandComponentGql } from "../command/CommandComponentGql";
import { OutputComponentGql } from "../output/OutputComponentGql";

const fragmentDefinition = graphql(`
  fragment TerminalEntryComponentGql on TerminalEntry2 {
    ... on TerminalCommand2 {
      ...CommandComponentGql
    }
    ... on TerminalOutput2 {
      ...OutputComponentGql
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  animate: boolean;
}

export function TerminalEntryComponentGql(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  switch (fragment.__typename) {
    case "TerminalCommand2":
      return (
        <CommandComponentGql fragment={fragment} animate={props.animate} />
      );
    case "TerminalOutput2":
      return <OutputComponentGql fragment={fragment} />;
  }
}
