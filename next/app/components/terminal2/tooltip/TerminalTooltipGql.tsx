import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { TerminalTooltip } from "./TerminalTooltip";

const fragmentDefinition = graphql(`
  fragment TerminalTooltipGql on TerminalTooltip2 {
    markdownBody
    timing
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function TerminalTooltipGql(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return <TerminalTooltip markdownBody={fragment.markdownBody} />;
}
