import React, { Component } from 'react';
import { findDOMNode } from 'react-dom';

class PayjpCheckout extends Component {

  constructor(props: Object) {
    super(props);
    this.windowAlertBackUp = window.alert;
    this.script = document.createElement('script');
    this.payjpCheckoutRef = null;
    this.script.setAttribute('src', 'https://checkout.pay.jp/');
    // @ts-ignore
    this.script.setAttribute('class', this.props.className);
    // @ts-ignore
    this.script.setAttribute('data-key', this.props.dataKey);
    // @ts-ignore
    this.props.dataPayjp && this.script.setAttribute('data-payjp', this.props.dataPayjp);
    // @ts-ignore
    this.props.dataPartial && this.script.setAttribute('data-partial', this.props.dataPartial);
    // @ts-ignore
    this.props.dataText && this.script.setAttribute('data-text', this.props.dataText);
    // @ts-ignore
    this.props.dataSubmitText && this.script.setAttribute('data-submit-text', this.props.dataSubmitText);
    // @ts-ignore
    this.props.dataTokenName && this.script.setAttribute('data-token-name', this.props.dataTokenName);
    // @ts-ignore
    this.props.dataPreviousToken && this.script.setAttribute('data-previous-token', this.props.dataPreviousToken);
    // @ts-ignore
    this.props.dataLang && this.script.setAttribute('data-lang', this.props.dataLang);
    this.script.setAttribute('data-on-created', 'reactPayjpCheckoutOnCreated');
    this.script.setAttribute('data-on-failed', 'reactPayjpCheckoutOnFailed');
    // @ts-ignore
    this.props.dataNamePlaceholder && this.script.setAttribute('data-name-placeholder', this.props.dataNamePlaceholder);
  }

  script: HTMLElement;
  payjpCheckoutRef: any;
  windowAlertBackUp: Function;

  static defaultProps = {
    className: 'payjp-button',
    dataKey: '',
    onCreatedHandler: () => {},
    onFailedHandler: () => {},
  };

  // @ts-ignore
  props: {
    onCreatedHandler?: Function;
    onFailedHandler?: Function;
    className: string;

    // =====================================================
    // Settings by PAY.JP
    // see PAY.JP docs for more info:
    //   https://pay.jp/docs/checkout
    // =====================================================

    // PAY.JP のパブリックキー
    dataKey: string;

    // ==============
    // Optional Props
    // ==============

    // PAY.JPで発行したOAuth Client IDを入れると、カード番号を入力することなく決済が可能なPAY ID決済を利用できるようになる
    // dataPayjp?: ?string;
    // trueとセットすると、カード情報フォーム入力後に自動的に送信(自動的なトークン作成)しない
    // dataPartial?: ?string;
    // カード情報入力フォームを開くのボタンのテキスト 	data-partialが false の時「カードで支払う」、true の時 「カード情報を送信する」
    // dataText?: ?string;
    // カード情報入力フォーム内の送信ボタンのテキスト
    // data-partialが false の時「カードで支払う」、true の時「カード情報を送信する」
    // dataSubmitText?: ?string;
    // 作成されたトークンがセットされるinput name(hidden)
    // dataTokenName?: ?string;
    // エラーにより画面に戻ってきた場合などに、再入力をさせないために作成済みのトークンを入れるパラメーター
    // dataPreviousToken?: ?string;
    // メッセージなどの表示言語
    // dataLang?: null | 'ja' | 'en';
    // カード名義のplaceholderにセットされる値
    // dataNamePlaceholder?: ?string;
  };

  componentWillMount() {
    // @ts-ignore
    window.reactPayjpCheckoutOnCreated = this.onCreated;
    // @ts-ignore
    window.reactPayjpCheckoutOnFailed = this.onFailed;
    // @ts-ignore
    window.reactPayjpCheckoutContext = this;
    // カード情報が不正のときに window.alert が payjp の checkout から呼ばれるため
    window.alert = () => {};
  }

  componentDidMount() {
    this.payjpCheckoutRef = findDOMNode(this.refs.payjpCheckoutRef);
    this.payjpCheckoutRef && this.payjpCheckoutRef.appendChild(this.script);
  }

  componentWillUnmount() {
    // すでに https://checkout.pay.jp/ の checkout.js が実行済みで、script タグを削除しているだけ
    this.payjpCheckoutRef.removeChild(this.script);
    // @ts-ignore
    window.reactPayjpCheckoutOnCreated = null;
    // @ts-ignore
    window.reactPayjpCheckoutOnFailed = null;
    // @ts-ignore
    window.reactPayjpCheckoutContext = null;
    // @ts-ignore
    window.alert = this.windowAlertBackUp;
    // @ts-ignore
    window.PayjpCheckout = null;
  }

  shouldComponentUpdate() {
    return false;
  }

  onCreated(response: any): void {
    const payload = { token: response.id }
    // @ts-ignore
    window.reactPayjpCheckoutContext.props.onCreatedHandler(payload);
  }

  onFailed(statusCode: string, errorResponse: Object): void {
    // @ts-ignore
    const payload = { message: errorResponse.message };
    // @ts-ignore
    window.reactPayjpCheckoutContext.props.onFailedHandler(payload);
  }

  render() {
    return <div ref="payjpCheckoutRef"></div>;
  }
}

export default PayjpCheckout;