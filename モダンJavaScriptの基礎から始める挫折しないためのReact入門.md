- DOMと仮想DOM
  - DOMはHTMLを機構造で操作するため、仮想DOMと比較するとレンダリングコストが高く、コードが複雑。
  - 仮想DOMはいきなりDOMを変更するのではなく、変更したら、まずは仮想DOMに変更を加えて、差分だけをDOMに反映する。


- npm/yarn
  - NPMに上がっているパッケージをnpm/yarnコマンドでローカルにダウンロード。NPMはパッケージのネット上で保管されている保管庫みたいなもの。
  - package.jsonにどのパッケージをどのバージョンでインストールするかを記述する。RailsのGemfileみたいなもの
  - package-lock.jsonまたはyarn.lockはRailsでいうGemfile.lock。パッケージをインストールした結果が反映される。
  - node_modulesが各モジュールの実態。サイズが膨大なため、GIthubには上げない。

- ECMAScriptとはJavascriptの書き方のルールみたいなもの。ES6(ES2015)から以下が使用できるようになった。
  - letやconst
  - スプレッド構文
  - Promise

- モジュールバンドラー。複数のJSやCSS、画像などを1つにまとめるためのもの。開発する際にはプログラムを分割して、使い回せるようにして、本番環境で動かす際には1つにまとめた方が読み込む際のパフォーマンスが良いため。
  - Webpack
  - Vite

- トランスパイラ。新しいJSの記法を古い記法に変換してくれるもの。ES2015で書いたものの、古いブラウザ（例えばIE）がその記法に対応しておらず、動かないことがあるため、トランスパイラで古いブラウザでもJSが動くようにするもの。
  - BABEL
  - SWC


- JSもRubyみたいに三項演算子が使える。書き方も同じ。A === B ? 'hoge' : 'fuga'みたいに書く。

- const、let、varの違い
  - varは変数の再宣言、上書きが可能
  - letは変数の再宣言は不可、上書きが可能
  - constは変数の再宣言および上書きは不可。ただし、オブジェクトのプロパティの変更、追加、削除などは可能。

- 分割代入はオブジェクトと配列で使用できる。
  - オブジェクトの分割代入は以下の通り。

      const myProfile = {
        name: "じゃけぇ",
        age: 31
      };
      const { name, age } = myProfile;
      const message2 = `名前は${name}です。年齢は${age}です。`;

  - 配列の分割代入は以下の通り。配列はオブジェクトと異なり、好きな変数名(nameやage)を割り当てることができる。
      const myProfile = ["じゃけぇ", 31];
      const [name, age] = myProfile;
      const message4 = `名前は${name}です。年齢は${age}です。`;

- オブジェクトの省略記法。オブジェクトのプロパティ名と値が同じ名前の時は値を省略できる。

  ```
    const name = "じゃけぇ";
    const age = 31;
    const myProfile = {
      name: name,
      age: age
    };
  ```

    上記のコードを以下のように省略することができる。
  ```
    const myProfile = {
      name,
      age
    };
    console.log(myProfile);
  ```

- スプレッド構文を使用してコピーすると参照が引き継がれない。

  以下のように書くとコピー元のarr4も値が変わってしまう。
  const arr4 = [10, 20];
  const arr8 = arr4;
  arr8[0] = 100;
  console.log(arr8);

  以下のように書くとコピー元のarr4の値は変わらない。
  const arr4 = [10, 20];
  const arr6 = [...arr4];
  arr6[0] = 100;
  console.log(arr4);

- mapで添字を使用したい場合は以下のように引数を増やす。
  const nameArr = ['田中', '山田', 'じゃけぇ'];
  nameArr.map((name, index) => console.log(`${index + 1}番目は${name}です`));

- filterは配列の値を欲しいものだけに絞り込むことができる。戻りは配列。
  const numArr = [1, 2, 3, 4, 5];
  const newNumArr = numArr.filter((num) => {
    return num % 2 === 1;
  });
  console.log(newNameArr);


- React
  以下コードのStrictModeのように記述するとReactの構文を厳密にチェックする。

  ```
  import { StrictMode } from "react";
  import { createRoot } from "react-dom/client";
  import { App } from "./App";

  const rootElement = document.getElementById("root");

  const root = createRoot(rootElement);

  root.render(
    <StrictMode>
      <App />
    </StrictMode>
  );
  ```


  - ReactでonClickイベントを定義するには{}で囲って、その中で関数を定義する。また、CSSをあてる際にもstyle={}の中で定義する。
  ```
  export const App = () => {
    const onclickButton = () => alert();
    const contentStyle = {
      color: "blue",
      fontSize: "18px",
      margin: 100,
    };
    return (
      <>
        <h1>こんちは</h1>
        <p style={contentStyle}>お元気ですか?</p>
        <button onClick={onclickButton}>OK</button>
      </>
    );
  };

  ```

  - Vue.jsのようにReactでもpropsを使用して、親から子供のコンポーネントに値を渡せる。以下コードではcolor="blue"と「こんちは」をColorfulMessageコンポーネントに渡している。

  ```
    import { ColorfulMessage } from "./components/ColorfulMessage";
    export const App = () => {
      const onclickButton = () => alert();
      return (
        <>
          <ColorfulMessage color="blue">こんちは</ColorfulMessage>
          <ColorfulMessage color="green"> こんちは！</ColorfulMessage>
          <button onClick={onclickButton}>OK</button>
        </>
      );
    };

  ```

  以下、子供のコンポーネントのpropsを使用して、分割代入して、color, childrenで親から受け取った値を使用できる。
  また、color,はオブジェクトの省略記法を使っている。color: color,と書かなくても良い。<ColorfulMessage color="green"> こんちは！</ColorfulMessage>のようにコンポーネントで囲まれた文字はprops.childrenで取得できる。

  ```
    export const ColorfulMessage = (props) => {
    const { color, children } = props;
    const contentStyleA = {
      color,
      fontSize: "18px",
    };
      return <p style={contentStyleA}>{children}</p>;
    };

  ```


  - useStateを使用することでコンポーネントの状態を管理できる。useStateはコンポーネントの中の最上位にしか定義できない。関数の中では定義できない。
    setNumでstateを更新する。

  ```
    import { useState } from "react";
    import { ColorfulMessage } from "./components/ColorfulMessage";
    export const App = () => {
      const [num, setNum] = useState(0);
      const onclickButton = () => setNum((prev) => prev + 1);
      return (
        <>
          <ColorfulMessage color="blue">こんちは</ColorfulMessage>
          <ColorfulMessage color="green"> こんちは！</ColorfulMessage>
          <button onClick={onclickButton}>OK</button>
          <p>{num}</p>
        </>
      );
    };
  ```

  - 再レンダリングされるのは以下3つのタイミング
    - stateが変更された時。
    - 親コンポーネントが変更された時。
    - propsに渡される値が変わった時。

  - useEffectは指定した変数の値が変更した時のみ中で定義したロジックが動く。

    ```
      const [num, setNum] = useState(0);
      useEffect (() => {
        console.log("--useEffect--");
      }, [num]);
    ```

  - named exportをインポートする際には指定した名前を指定しなければならない。一方、default exportはインポートした先で自分で名前を定義できる。

