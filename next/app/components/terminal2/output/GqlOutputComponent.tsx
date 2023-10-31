import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { OutputComponent } from "./OutputComponent";

const fragmentDefinition = graphql(`
  fragment GqlOutputComponent on TerminalOutput2 {
    output
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlOutputComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return <OutputComponent output={fragment.output} />;
}
