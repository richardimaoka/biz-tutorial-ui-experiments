import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlColumnTab.module.css";
import { GqlColumnTabIcon } from "./GqlColumnTabIcon";
import { LinkSearchParams } from "@/app/components/link/LinkSearchParams";

const fragmentDefinition = graphql(`
  fragment GqlColumnTab on ColumnWrapper2 {
    columnName
    columnDisplayName
    ...GqlColumnTabIcon
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  isSelected?: boolean;
}

export function GqlColumnTab(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  console.log(fragment);

  // CSS style for the outer component
  const selectStyle = props.isSelected ? styles.selected : styles.unselected;
  const outerClassName = `${styles.component} ${selectStyle}`;

  // Display name of this tab
  const displayName = fragment.columnDisplayName
    ? fragment.columnDisplayName
    : fragment.columnName;

  // Search params (a.k.a. query params) for the link (browser navigation)
  const searchParams = fragment.columnName
    ? {
        column: fragment.columnName,
      }
    : ({} as Record<string, string>);

  return (
    <LinkSearchParams searchParams={searchParams}>
      <div className={outerClassName}>
        <span className={styles.smartphone}>
          <GqlColumnTabIcon fragment={fragment} />
        </span>
        <span className={styles.desktop}>{displayName}</span>
      </div>
    </LinkSearchParams>
  );
}
