import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import axios, { AxiosRequestConfig } from "axios";
import styles from "./main.module.scss"; // SCSSモジュールのインポート

type Condition = {
  condition_name: string;
  // start_date: string;
  // end_date: string;
  damage_point: number;
};
export default function Main() {
  const [conditions, setConditions] = useState<Condition[]>([]);
  const [allConditions, setAllConditions] = useState<Condition[]>([]);
  const [errorMessage, setErrorMessage] = useState("");
  const router = useRouter();
  const [genkiHP, setGenkiHP] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/login"); // トークンがなければログインページにリダイレクト
    } else {
      fetchConditionsDisplay(token);
      todayPoint(token);
      handlefetchConditions(token);
    }
  }, [router]);

  const fetchConditionsDisplay = async (token: string) => {
    const url = "http://localhost:8080";
    console.log("Fetching conditions...");
    const options: AxiosRequestConfig = {
      url: `${url}/users/me/condition/today`,
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
        const data = response.data;
        setConditions(data);
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

  const todayPoint = async (token: string) => {
    const url = "http://localhost:8080";
    console.log("Fetching today's point...");
    const options: AxiosRequestConfig = {
      url: `${url}/users/me/condition/today/point`,
      method: "GET",
      headers: {
        "x-token": token,
      },
    };

    try {
      const response = await axios(options);
      const genkiHP = response.data;
      console.log(`Today's Genki HP:`, genkiHP);
      setGenkiHP(genkiHP); // コメントを解除して状態を更新
      setErrorMessage("");
    } catch (error) {
      console.error("Error fetching today's point:", error);
      setErrorMessage(
        "情報の取得中にサーバーでエラーが発生しました。しばらくしてから再度試してください。"
      );
    }
  };

  const handlefetchConditions = async (token: string) => {
    const url = "http://localhost:8080";
    console.log("Fetching conditions...");
    const options: AxiosRequestConfig = {
      url: `${url}/users/me/condition`,
      method: "GET",
      headers: {
        "x-token": token,
      },
    };

    try {
      const response = await axios(options);
      const data = response.data;
      console.log("allConditions:", data);
      setAllConditions(data);
      setErrorMessage("");
    } catch (error) {
      console.error("Error fetching conditions:", error);
      setErrorMessage(
        "情報の取得中にサーバーでエラーが発生しました。しばらくしてから再度試してください。"
      );
    }
  };

  const handleConditionClick = () => {
    router.push("/condition"); // コンディションページへのリダイレクト
  };

  return (
    <div className={styles.mainContainer}>
      <div className={styles.card}>
        <h1>今日の元気予報</h1>
        <h2>
          {genkiHP !== null && <p>{genkiHP}/100</p>} {/* 元気ポイントの表示 */}
        </h2>
        {/* Rest of your code */}
        <img
          src="/girl1.png"
          alt="Description of image"
          className={styles.cardImage}
        />
      </div>
      <div className={styles.cards}>
      <div className={styles.cardmini}>
        <h2>予報詳細：</h2>
        <div className={styles.detail}>
          <ul>
            {conditions.map((condition, index) => (
              <li key={index}>
                {condition.condition_name}で{condition.damage_point}pt
              </li>
            ))}
          </ul>
        </div>
        </div>
        <div className={styles.cardmini}>
        <h2>登録済みの体調</h2>
        <div className={styles.detail}>
        <ul>
          {allConditions.map((allCondition, index) => (
            <li key={index}>
              {allCondition.Name}で-{allCondition.DamagePoint}pt
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
    </div>
  );
}
