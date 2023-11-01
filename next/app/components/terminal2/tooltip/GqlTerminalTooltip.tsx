import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { TerminalTooltip } from "./TerminalTooltip";
import { TerminalTooltipDelay } from "./TerminalTooltipDelay";

const fragmentDefinition = graphql(`
  fragment GqlTerminalTooltip on TerminalTooltip2 {
    markdownBody
    timing
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlTerminalTooltip(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const delay = fragment.timing === "END";
  return (
    <TerminalTooltipDelay delay={delay}>
      <TerminalTooltip markdownBody={fragment.markdownBody} />
    </TerminalTooltipDelay>
  );
}
