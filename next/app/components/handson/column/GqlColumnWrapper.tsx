import { GqlBrowserColumn } from "../../browser/GqlBrowserColumn";
import { GqlModalComponent } from "../../modal/GqlModalComponent";
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
    modal {
      ...GqlModalComponent
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

  const modal = fragment.modal;

  switch (column.__typename) {
    case "TerminalColumn":
      return (
        // This <div> is needed (1) to set the common background color,
        // and (2) to set `position: relative` to hold common `position: absolute` items (e.g.) navigation buttons
        <div className={styles.component}>
          <GqlTerminalColumn fragment={column} selectIndex={0} />
          {modal && <GqlModalComponent fragment={modal} />}
        </div>
      );
    case "SourceCodeColumn":
      return (
        // This <div> is needed (1) to set the common background color,
        // and (2) to set `position: relative` to hold common `position: absolute` items (e.g.) navigation buttons
        <div className={styles.component}>
          <GqlSourceCodeColumn
            fragment={column}
            defaultFocusColumn={props.defaultFocusColumn}
          />
          {modal && <GqlModalComponent fragment={modal} />}
        </div>
      );
    case "BrowserColumn":
      return (
        // This <div> is needed (1) to set the common background color,
        // and (2) to set `position: relative` to hold common `position: absolute` items (e.g.) navigation buttons
        <div className={styles.component}>
          <GqlBrowserColumn fragment={column} />
          {modal && <GqlModalComponent fragment={modal} />}
        </div>
      );
  }
}
