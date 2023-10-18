import { Tooltip } from "../../tooltip/Tooltip";

interface Props {
  markdownBody: string;
  hidden?: boolean;
}

export type TerminalTooltipProps = Props;

export function TerminalTooltip(props: Props) {
  return <Tooltip markdownBody={props.markdownBody} />;
}
