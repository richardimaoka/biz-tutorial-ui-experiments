import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { GqlBrowser } from "./GqlBrowser";
import styles from "./GqlBrowserColumn.module.css";

const fragmentDefinition = graphql(`
  fragment GqlBrowserColumn on BrowserColumn {
    browser {
      ...GqlBrowser
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlBrowserColumn(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      <GqlBrowser fragment={fragment.browser} />
    </div>
  );
}
