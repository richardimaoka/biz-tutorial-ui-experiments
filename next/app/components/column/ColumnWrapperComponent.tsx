import { FragmentType, graphql, useFragment } from "@/libs/gql";

import styles from "./style.module.css";
import { SourceCodeColumn } from "../sourcecode/SourceCodeColumn";
import { TerminalColumn } from "../terminal/TerminalColumn";
import { useState } from "react";

const fragmentDefinition = graphql(`
  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {
    name
    column {
      __typename
      ... on SourceCodeColumn {
        ...SourceCodeColumn_Fragment
      }

      ... on TerminalColumn {
        ...TerminalColumn_Fragment
      }
    }
  }
`);

interface ColumnWrapperComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  step: string;
  skipAnimation?: boolean;
}

export const ColumnWrapperComponent = (
  props: ColumnWrapperComponentProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment?.column?.__typename) {
    return <div></div>;
  }

  switch (fragment.column.__typename) {
    case "SourceCodeColumn":
      return <SourceCodeColumn fragment={fragment.column} step={props.step} />;
    case "TerminalColumn":
      return (
        <TerminalColumn
          fragment={fragment.column}
          skipAnimation={props.skipAnimation}
        />
      );
    default:
      return <div>default</div>;
  }
};
