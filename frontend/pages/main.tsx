import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import axios, { AxiosRequestConfig } from "axios";
import { format, isWithinInterval, parseISO } from "date-fns";
import styles from "./main.module.scss"; // SCSSモジュールのインポート

type Condition = {
  condition_name: string;
  start_date: string;
  end_date: string;
  damage_point: number;
};
export default function Main() {
  const [conditions, setConditions] = useState<Condition[]>([]);
  const [errorMessage, setErrorMessage] = useState("");
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/login"); // トークンがなければログインページにリダイレクト
    } else {
      fetchConditions(token);
    }
  }, [router]);

  const fetchConditions = async (token: string) => {
    const url = "http://localhost:8080";
    console.log("Fetching conditions...");
    const options: AxiosRequestConfig = {
      url: `${url}/users/me/condition`,
      method: "GET",
      headers: {
        "x-token": token,
      },
    };
    console.log("Options:", options);

    // エラー処理
    try {
      const response = await axios(options);
      if (Array.isArray(response.data)) {
        const today = new Date();
        const filteredConditions = response.data.filter(
          (condition: Condition) =>
            isWithinInterval(today, {
              start: parseISO(condition.start_date),
              end: parseISO(condition.end_date),
            })
        );
        setConditions(filteredConditions);
      } else {
        throw new TypeError("Received data is not an array");
      }
      setErrorMessage("");
    } catch (error) {
      console.error("Error fetching conditions:", error);
      setErrorMessage(
        "情報の取得中にサーバーでエラーが発生しました。しばらくしてから再度試してください。"
      );
    }
    console.log("Conditions after fetching:", conditions); // ログに状態を出力
  };

  const handleConditionClick = () => {
    router.push("/condition"); // コンディションページへのリダイレクト
  };

  return (
    <div className={styles.mainContainer}>
      <div className={styles.card}>
        <h1>今日の元気予報</h1>
        {errorMessage && <div className={styles.error}>{errorMessage}</div>}
      </div>
      <div className={styles.card}>
        <h2>予報詳細：</h2>
        <div className={styles.detail}>
          <ul>
            {conditions.map((condition, index) => (
              <li key={index}>
                {condition.condition_name}：{condition.damage_point}
              </li>
            ))}
          </ul>
        </div>
        <button
          className={styles.ConditionButton}
          onClick={handleConditionClick}
        >
          体調の新規登録
        </button>
      </div>
    </div>
  );
}
