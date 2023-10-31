import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { GqlCommandComponent } from "../command/GqlCommandComponent";
import { GqlOutputComponent } from "../output/GqlOutputComponent";

const fragmentDefinition = graphql(`
  fragment GqlTerminalEntryComponent on TerminalEntry2 {
    ... on TerminalCommand2 {
      ...GqlCommandComponent
    }
    ... on TerminalOutput2 {
      ...GqlOutputComponent
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  animate: boolean;
}

export function GqlTerminalEntryComponent(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  switch (fragment.__typename) {
    case "TerminalCommand2":
      return (
        <GqlCommandComponent fragment={fragment} animate={props.animate} />
      );
    case "TerminalOutput2":
      return <GqlOutputComponent fragment={fragment} />;
  }
}
