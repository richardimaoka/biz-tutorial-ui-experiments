import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { CommandComponent } from "../command/CommandComponent";
import { OutputComponent } from "../output/OutputComponent";

const fragmentDefinition = graphql(`
  fragment GqlTerminalEntryComponent on TerminalEntry {
    entryType
    text
    isCommandExecuted
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  isLastEntry: boolean;
}

export function GqlTerminalEntryComponent(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const beforeExecution = !fragment.isCommandExecuted;

  switch (fragment.entryType) {
    case "COMMAND":
      return (
        <CommandComponent
          command={fragment.text}
          animate={props.isLastEntry && beforeExecution}
        />
      );
    case "OUTPUT":
      return <OutputComponent output={fragment.text} />;
  }
}
