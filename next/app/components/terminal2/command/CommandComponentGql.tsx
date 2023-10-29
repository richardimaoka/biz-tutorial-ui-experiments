import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { CommandComponent } from "./CommandComponent";

const fragmentDefinition = graphql(`
  fragment CommandComponentGql on TerminalCommand2 {
    command
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  animate: boolean;
}

export function CommandComponentGql(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return (
    <CommandComponent command={fragment.command} animate={props.animate} />
  );
}
