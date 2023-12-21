import { GqlHandsonComponent } from "@/app/components/handson/GqlHandsonComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";
import { GqlNavigation } from "../components/navigation/GqlNavigation";
import styles from "./page.module.css";
import { GqlSlideshowComponent } from "../components/slideshow/GqlSlideshowComponent";
import { print } from "graphql";

const queryDefinition = graphql(`
  query appTutorialPage($tutorial: String!, $step: String) {
    page(tutorial: $tutorial, step: $step) {
      ...GqlHandsonComponent
      ...GqlNavigation
      ...GqlSlideshowComponent
      mode
    }
  }
`);

/**
 * page.js
 * https://nextjs.org/docs/app/api-reference/file-conventions/page
 *
 *
 */
interface PageParams {
  params: {
    tutorial: string;
  };
  searchParams: {
    step?: string;
  };
}

export default async function Page(props: PageParams) {
  // console.log("----------------------------------------");
  // console.log(print(queryDefinition));

  const gqlEndPoint = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
  if (typeof gqlEndPoint != "string") {
    throw new Error("Next.js server gave wrong GraphQL endpoint URL.");
  }

  const variables = {
    tutorial: props.params.tutorial,
    step: props.searchParams.step,
  };
  const data = await request(gqlEndPoint, queryDefinition, variables);

  const fragment = data.page;
  const mode = fragment?.mode === "SLIDESHOW" ? "SLIDESHOW" : "HANDSON";

  switch (mode) {
    case "SLIDESHOW":
      return (
        <div className={styles.component}>
          {fragment && <GqlSlideshowComponent fragment={fragment} />}
          {fragment && (
            <GqlNavigation
              tutorial={props.params.tutorial}
              fragment={fragment}
              toInitial
            />
          )}
        </div>
      );
    case "HANDSON":
      return (
        <div className={styles.component}>
          {fragment && <GqlHandsonComponent fragment={fragment} />}
          {fragment && (
            <GqlNavigation
              tutorial={props.params.tutorial}
              fragment={fragment}
            />
          )}
        </div>
      );
  }
}
