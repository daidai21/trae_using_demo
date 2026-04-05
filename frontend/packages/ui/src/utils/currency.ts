export enum Currency {
  CNY = 'CNY',
  USD = 'USD',
  IDR = 'IDR',
}

interface CurrencyConfig {
  symbol: string;
  decimalSeparator: string;
  thousandSeparator: string;
  decimalPlaces: number;
  symbolPosition: 'before' | 'after';
}

const currencyConfigs: Record<Currency, CurrencyConfig> = {
  [Currency.CNY]: {
    symbol: '¥',
    decimalSeparator: '.',
    thousandSeparator: ',',
    decimalPlaces: 2,
    symbolPosition: 'before',
  },
  [Currency.USD]: {
    symbol: '$',
    decimalSeparator: '.',
    thousandSeparator: ',',
    decimalPlaces: 2,
    symbolPosition: 'before',
  },
  [Currency.IDR]: {
    symbol: 'Rp',
    decimalSeparator: ',',
    thousandSeparator: '.',
    decimalPlaces: 0,
    symbolPosition: 'before',
  },
};

export const formatCurrency = (amount: number, currency: Currency = Currency.CNY): string => {
  const config = currencyConfigs[currency];
  
  const fixedAmount = amount.toFixed(config.decimalPlaces);
  const [integerPart, decimalPart] = fixedAmount.split('.');
  
  const formattedInteger = integerPart.replace(/\B(?=(\d{3})+(?!\d))/g, config.thousandSeparator);
  
  let result = formattedInteger;
  if (config.decimalPlaces > 0 && decimalPart) {
    result += config.decimalSeparator + decimalPart;
  }
  
  if (config.symbolPosition === 'before') {
    result = config.symbol + ' ' + result;
  } else {
    result = result + ' ' + config.symbol;
  }
  
  return result;
};

export const getCurrencySymbol = (currency: Currency): string => {
  return currencyConfigs[currency].symbol;
};
