import { GqlTerminalComponent } from "@/app/components/terminal2/GqlTerminalComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTerminalTooltipPage {
    _test {
      appTestTerminalTooltipPage {
        ...GqlTerminalComponent
      }
    }
  }
`);
//https://github.com/richardimaoka/biz-tutorial-ui-experiments/commit/86255d7817b40d1455888b22b75a50a05028f620
export default async function Page() {
  const data = await request("http://localhost:8080/query", queryDefinition);

  const fragment = data._test?.appTestTerminalTooltipPage;

  if (!fragment) {
    return <>no data</>;
  }

  return (
    <div style={{ height: "95svh" }}>
      <GqlTerminalComponent fragment={fragment} />
    </div>
  );
}
