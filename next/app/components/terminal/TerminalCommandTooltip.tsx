import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./TerminalCommandTooltip.module.css";

const fragmentDefinition = graphql(`
  fragment TerminalCommandTooltip on TerminalCommand {
    tooltip
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function TerminalCommandTooltip(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.tooltip) {
    return <></>;
  }

  return <div className={styles.component}>{fragment.tooltip}</div>;
}
