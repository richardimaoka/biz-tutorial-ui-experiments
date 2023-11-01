import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { CommandComponent } from "../command/CommandComponent";
import { OutputComponent } from "../output/OutputComponent";

const fragmentDefinition = graphql(`
  fragment GqlTerminalEntryComponent on TerminalEntry {
    entryType
    text
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  animate: boolean;
  completedCallback?: () => void;
}

export function GqlTerminalEntryComponent(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  switch (fragment.entryType) {
    case "COMMAND":
      return (
        <CommandComponent
          command={fragment.text}
          animate={props.animate}
          completedCallback={props.completedCallback}
        />
      );
    case "OUTPUT":
      return <OutputComponent output={fragment.text} />;
  }
}
