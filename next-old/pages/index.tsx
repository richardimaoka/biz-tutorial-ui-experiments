import { GetServerSideProps } from "next";
import { ParsedUrlQuery } from "querystring";
import { PageColumns2 } from "../components/column/PageColumns2";
import { Navigation } from "../components/navigation/Navigation";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";
import { queryParamToString } from "../libs/queryString";
// import { print } from "graphql";

const queryDefinition = graphql(/* GraphQL */ `
  query IndexSsrPage($tutorial: String!, $step: String, $openFilePath: String) {
    page(tutorial: $tutorial, step: $step) {
      __typename
      ...NavigationFragment
      ...PageColumns2Fragment
    }
  }
`);

interface PageParams extends ParsedUrlQuery {
  step: string;
}

export const getServerSideProps: GetServerSideProps<
  IndexSsrPageQuery,
  PageParams
> = async (context) => {
  console.log(
    "getServerSideProps called for index.tsx",
    new Date().toUTCString()
  );
  const step = queryParamToString(context.query.step);
  const openFilePath = queryParamToString(context.query.openFilePath);

  const { data } = await client.query({
    query: queryDefinition,
    variables: {
      tutorial: "sign-in-with-google",
      step: step,
      openFilePath: openFilePath,
    },
  });

  return {
    props: data,
  };
};

export default function Home({ page }: IndexSsrPageQuery) {
  // const router = useRouter();

  // to debug query
  //console.log(print(queryDefinition));

  // const terminals = currentPage?.terminals;
  // const [currentTerminalIndex] = useState(0);
  // const currentTerminal = terminals && terminals[currentTerminalIndex];
  // const currentDirectory = currentTerminal?.currentDirectory
  //   ? currentTerminal.currentDirectory
  //   : undefined;

  // useEffect(() => {
  //   const handleKeyDown = (event: KeyboardEvent) => {
  //     if (event.code === "Space") {
  //       if (event.shiftKey) {
  //         prevStep &&
  //           router.push(`/${constructQueryString(prevStep, openFilePath)}`);
  //       } else {
  //         nextStep &&
  //           router.push(`/${constructQueryString(nextStep, openFilePath)}`);
  //       }
  //     }
  //   };
  //   document.addEventListener("keydown", handleKeyDown);

  //   // Don't forget to clean up
  //   return function cleanup() {
  //     document.removeEventListener("keydown", handleKeyDown);
  //   };
  // }, [router, nextStep, prevStep, openFilePath]);

  return (
    <>
      {page && <PageColumns2 fragment={page} />}
      {page && <Navigation fragment={page} />}
    </>
  );
}
