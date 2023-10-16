import { Column } from "./column/Column";
import styles from "./Columns.module.css";

interface Props {}

export async function Columns(props: Props) {
  const arr = [1, 2, 3, 4, 5, 6, 7, 8];
  return (
    <div className={styles.component}>
      {arr.map((x, index) => (
        <div key={index} className={styles.eachColumn}>
          <Column>
            <div
              style={{
                height: "100%",
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                color: "white",
                fontSize: "80px",
              }}
            >
              {index}
            </div>
          </Column>
        </div>
      ))}
    </div>
  );
}
