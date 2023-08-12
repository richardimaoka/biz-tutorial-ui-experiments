import Image from "next/image";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment DevToolsColumn_Fragment on DevToolsColumn {
    width
    height
    path
  }
`);

export interface DevToolsColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const DevToolsColumn = (props: DevToolsColumnProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.path || !fragment.width || !fragment.height) {
    return <></>;
  }

  return (
    <Image
      className={styles.image}
      src={fragment.path}
      width={fragment.width}
      height={fragment.height}
      alt="DevTools screenshot"
    />
  );
};
