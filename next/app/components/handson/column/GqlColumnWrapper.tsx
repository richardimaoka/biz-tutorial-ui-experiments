import { GqlBrowserColumn } from "../../browser/GqlBrowserColumn";
import { GqlSourceCodeColumn } from "../../sourcecode/GqlSourceCodeColumn";
import { GqlTerminalColumn } from "../../terminal/GqlTerminalColumn";
import styles from "./GqlColumnWrapper.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlColumnWrapper on ColumnWrapper {
    columnName
    column {
      # if you forget this, the resulting fragment will have __typename = undefined
      __typename
      #
      # for each column type
      #
      ... on TerminalColumn {
        ...GqlTerminalColumn
      }
      ... on SourceCodeColumn {
        ...GqlSourceCodeColumn
      }
      ... on BrowserColumn {
        ...GqlBrowserColumn
      }
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  defaultFocusColumn?: string;
}

export function GqlColumnWrapper(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const column = fragment.column;

  if (!column.__typename) {
    throw new Error(
      "__typename got undefined - define __typename in GraphQL fragment/query"
    );
  }

  switch (column.__typename) {
    case "TerminalColumn":
      return (
        <div className={styles.component}>
          <GqlTerminalColumn fragment={column} selectIndex={0} />
        </div>
      );
    case "SourceCodeColumn":
      return (
        <div className={styles.component}>
          <GqlSourceCodeColumn
            fragment={column}
            defaultFocusColumn={props.defaultFocusColumn}
          />
        </div>
      );
    case "BrowserColumn":
      return (
        <div className={styles.component}>
          <GqlBrowserColumn fragment={column} />
        </div>
      );
  }
}
