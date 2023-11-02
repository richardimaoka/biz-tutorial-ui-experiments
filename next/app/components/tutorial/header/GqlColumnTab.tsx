"use client";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { useRouter } from "next/navigation";
import styles from "./GqlColumnTab.module.css";
import { GqlColumnTabIcon } from "./GqlColumnTabIcon";

const fragmentDefinition = graphql(`
  fragment GqlColumnTab on ColumnWrapper2 {
    columnName
    ...GqlColumnTabIcon
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  isSelected?: boolean;
}

export function GqlColumnTab(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const router = useRouter();

  // TODO: href calculation
  const href = "/";

  function onClick() {
    // need to use router.replace instaed of <Link> not to mess up the browser history
    router.replace(href);
  }

  const selectStyle = props.isSelected ? styles.selected : styles.unselected;
  const outerClassName = `${styles.component} ${selectStyle}`;

  return (
    <button className={outerClassName} onClick={onClick}>
      <span className={styles.smartphone}>
        <GqlColumnTabIcon fragment={fragment} />
      </span>
      <span className={styles.desktop}>{fragment.columnName}</span>
    </button>
  );
}
