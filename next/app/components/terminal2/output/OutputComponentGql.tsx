import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { OutputComponent } from "./OutputComponent";

const fragmentDefinition = graphql(`
  fragment OutputComponentGql on TerminalOutput2 {
    output
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function OutputComponentGql(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return <OutputComponent output={fragment.output} />;
}
