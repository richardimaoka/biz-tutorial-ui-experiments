import Link from "next/link";
import { ChromeIcon } from "../icons/ChromeIcon";
import { FileLinesIcon } from "../icons/FileLinesIcon";
import { SourceCodeIcon } from "../icons/SourceCodeIcon";
import { TerminalIcon } from "../icons/TerminalIcon";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment ColumnTab_Fragment on ColumnWrapper {
    name
    column {
      __typename
    }
  }
`);

export interface ColumnTabProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  isSelected: boolean;
  openFilePath?: string;
  step: string;
}

export const ColumnTab = (props: ColumnTabProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const selectStyle = props.isSelected ? styles.selected : styles.unselected;

  const Icon = (): JSX.Element => {
    if (!fragment?.column?.__typename) {
      return <FileLinesIcon />;
    }

    switch (fragment.column.__typename) {
      case "SourceCodeColumn":
        return <SourceCodeIcon />;
      case "TerminalColumn":
        return <TerminalIcon />;
      case "BackgroundImageColumn":
        return <ChromeIcon />;
      case "ImageDescriptionColumn":
        return <FileLinesIcon />;
      case "MarkdownColumn":
        return <FileLinesIcon />;
      case "BrowserColumn":
        return <ChromeIcon />;
      case "DevToolsColumn":
        return <ChromeIcon />;
    }
  };

  return (
    <div className={`${styles.tab} ${selectStyle}`}>
      <Link
        href={{
          query: {
            column: encodeURIComponent(fragment.name ? fragment.name : ""),
            openFilePath: props.openFilePath,
            step: props.step,
          },
        }}
      >
        <span className={styles.smartphone}>
          <Icon />
        </span>
        <span className={styles.desktop}>{fragment.name}</span>
      </Link>
    </div>
  );
};
