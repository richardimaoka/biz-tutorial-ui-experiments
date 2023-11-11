"use client";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlColumnTab.module.css";
import { GqlColumnTabIcon } from "./GqlColumnTabIcon";
import { LinkSearchParams } from "@/app/components/link/LinkSearchParams";
import { useSearchParams } from "next/navigation";

const fragmentDefinition = graphql(`
  fragment GqlColumnTab on ColumnWrapper {
    columnName
    columnDisplayName
    ...GqlColumnTabIcon
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  tabIndex: number;
}

export function GqlColumnTab(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const searchParams = useSearchParams();

  // CSS style for the outer component
  const columnParam = searchParams.get("column");
  const isSelected = columnParam
    ? columnParam === fragment.columnName
    : props.tabIndex === 0;
  const selectStyle = isSelected ? styles.selected : styles.unselected;
  const outerClassName = `${styles.component} ${selectStyle}`;

  // Display name of this tab
  const displayName = fragment.columnDisplayName
    ? fragment.columnDisplayName
    : fragment.columnName;

  // Search params (a.k.a. query params) for the link (browser navigation)
  const newSearchParams = fragment.columnName
    ? {
        column: fragment.columnName,
      }
    : ({} as Record<string, string>);

  return (
    <LinkSearchParams searchParams={newSearchParams}>
      <div className={outerClassName}>
        <span className={styles.smartphone}>
          <GqlColumnTabIcon fragment={fragment} />
        </span>
        <span className={styles.desktop}>{displayName}</span>
      </div>
    </LinkSearchParams>
  );
}
